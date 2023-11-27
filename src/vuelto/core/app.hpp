#pragma once
#include <GLFW/glfw3.h>

#include "../tools/definitions.h"

namespace Vuelto {

class Window;

namespace Application {

void Init(const bool softwareRenderer = false);
Window CreateWindow(int width, int height, const char *title, bool resizable);
void DestroyWindow(Window win);
void Terminate();

}  // namespace Application

class Window {
 public:
  GLFWwindow *window;
  int width;
  int height;
  const char *title;
  bool WindowShouldClose();
  void WindowRefresh();
};

}  // namespace Vuelto
