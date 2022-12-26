package vkng_surface_sdl2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/cockroachdb/errors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/extensions/v2/khr_surface"
	khr_surface_driver "github.com/vkngwrapper/extensions/v2/khr_surface/driver"
	_ "github.com/vkngwrapper/integrations/sdl2/v2/vulkan"
	"unsafe"
)

// CreateSurface generates a khr_surface.Surface from an SDL2 Window
//
// instance - The Instance that will own the new Surface
//
// extension - A khr_surface.Extension object
//
// window - And SDL2 Window that the new khr_surface.Surface will represent
func CreateSurface(instance core1_0.Instance, extension khr_surface.Extension, window *sdl.Window) (khr_surface.Surface, error) {
	if instance == nil {
		panic("instance cannot be nil")
	}

	surfacePtrUnsafe, err := window.VulkanCreateSurface((*C.VkInstance)(unsafe.Pointer(instance.Handle())))
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve surface from SDL")
	}

	surfaceHandlePtr := (*khr_surface_driver.VkSurfaceKHR)(surfacePtrUnsafe)

	return extension.CreateSurfaceFromHandle(*surfaceHandlePtr)
}
