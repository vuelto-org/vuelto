#pragma once

#include <GLFW/glfw3.h>

#include "../tools/definitions.hpp"
#include "renderer.hpp"

namespace Vuelto {

class Window;

namespace Application {

void Init();
void InitMultipleWindows();
Window CreateWindow(int width, int height, const char *title, bool resizable);
Vuelto::Renderer2D CreateRenderer2D(Window win);
Vuelto::SoftwareRenderer CreateSoftwareRenderer(Window win);

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
