#include "app.hpp"

#include <iostream>

#include "../tools/definitions.h"
#include "renderer.hpp"

namespace Vuelto {

void Init() {
  if (!glfwInit()) {
    std::cout << "GLFW Init failed\n";
  }
}

void Terminate() { glfwTerminate(); }

namespace Application {

void framebuffer_size_callback(GLFWwindow *window, int newWidth, int newHeight) {
  glViewport(0, 0, newWidth, newHeight);
  Vuelto::SoftwareRenderer::ResizeBuffer(newWidth, newHeight);
}

Window CreateWindow(int width, int height, const char *title, bool resizable) {
  glfwWindowHint(GLFW_RESIZABLE, resizable);

  GLFWwindow *glfw_window = glfwCreateWindow(width, height, title, NULL, NULL);
  if (!glfw_window) {
    glfwTerminate();
    std::cout << "GLFW Window creation failed\n";
  }

  glfwSetFramebufferSizeCallback(glfw_window, framebuffer_size_callback);

  glfwMakeContextCurrent(glfw_window);

  Window window;
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
  glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
}
}  // namespace Vuelto
