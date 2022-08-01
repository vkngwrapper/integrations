# Vkngwrapper Integrations

Each folder in this repository is a distinct module (in order to prevent all integrations from sharing
 dependencies) that adds utilities for integrating a specific heavyweight go package into your downstream
 libraries.

[go-sdl2](sdl2) - An integration for [Veandco's SDL2 wrapper](https://github.com/veandco/go-sdl2) that
 creates vkngwrapper khr_surface.Surface objects from sdl2.Window objects
