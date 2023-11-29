#pragma once
#include <GLFW/glfw3.h>

#include "../tools/definitions.hpp"

namespace Vuelto {

class Window;

namespace Application {

void Init(const bool softwareRenderer = false);
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
