package vkng_surface_sdl2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	"github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	_ "github.com/vkngwrapper/integrations/sdl2/v3/vulkan"
)

// CreateSurface generates a khr_surface.Surface from an SDL2 Window
//
// instance - The Instance that will own the new Surface
//
// extension - A khr_surface.Extension object
//
// window - And SDL2 Window that the new khr_surface.Surface will represent
func CreateSurface(instance core1_0.Instance, extension khr_surface.ExtensionDriver, window *sdl.Window) (khr_surface.Surface, error) {
	if !instance.Initialized() {
		panic("instance cannot be uninitialized")
	}

	surfacePtrUnsafe, err := window.VulkanCreateSurface((*C.VkInstance)(unsafe.Pointer(instance.Handle())))
	if err != nil {
		return khr_surface.Surface{}, errors.Wrap(err, "could not retrieve surface from SDL")
	}

	surfaceHandlePtr := (*khr_surface_loader.VkSurfaceKHR)(surfacePtrUnsafe)

	return extension.CreateSurfaceFromHandle(*surfaceHandlePtr)
}
