package vkng_surface_sdl2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/cockroachdb/errors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/extensions/khr_surface"
	khr_surface_driver "github.com/vkngwrapper/extensions/khr_surface/driver"
	"unsafe"
)

func CreateSurface(instance core1_0.Instance, extension khr_surface.Extension, window *sdl.Window) (khr_surface.Surface, error) {
	surfacePtrUnsafe, err := window.VulkanCreateSurface((*C.VkInstance)(unsafe.Pointer(instance.Handle())))
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve surface from SDL")
	}

	return extension.CreateSurfaceFromHandle((khr_surface_driver.VkSurfaceKHR)(surfacePtrUnsafe))
}
