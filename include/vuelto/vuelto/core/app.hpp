#pragma once

// clang-format off
#include "../vendor/GLAD/include/glad/glad.h"
#include <GLFW/glfw3.h>
// clang-format on

// #include "../tools/definitions.hpp"

namespace Vuelto {

class Window;

namespace Application {

void Init();
void InitMultipleWindows();
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
  void Refresh();
  void MakeContextCurrent();
};

}  // namespace Vuelto
