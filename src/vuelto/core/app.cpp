#include "app.hpp"

#include <iostream>

#include "../tools/definitions.h"

namespace Vuelto {

void Init() {
  if (!glfwInit()) {
    std::cout << "GLFW Init failed\n";
  }
}

void Terminate() { glfwTerminate(); }

namespace Application {

Window CreateWindow(int width, int height, const char *title) {
  GLFWwindow *glfw_window = glfwCreateWindow(width, height, title, NULL, NULL);
  if (!glfw_window) {
    glfwTerminate();
    std::cout << "GLFW Window creation failed\n";
  }

  glfwMakeContextCurrent(glfw_window);

  Window window;  // Allocate memory for the window object
  window.width = width;
  window.height = height;
  window.title = title;
  window.window = glfw_window;
  return window;
}

void DestroyWindow(Window win) { glfwDestroyWindow(win.window); }

}  // namespace Application

bool Window::WindowShouldClose() {
  Window::WindowRefresh();
  return glfwWindowShouldClose(window);
}

void Window::WindowRefresh() {
  glfwSwapBuffers(window);
  glfwPollEvents();
  glClear(GL_COLOR_BUFFER_BIT);
}
}  // namespace Vuelto
