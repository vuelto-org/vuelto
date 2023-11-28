#include "app.hpp"

#include <iostream>

#include "../tools/definitions.h"
#include "renderer.hpp"

namespace Vuelto {
namespace Application {

bool SoftwareRendererEnabled = false;

void Init(const bool softwareRenderer) {
  if (!glfwInit()) {
    std::cout << "GLFW Init failed\n";
  }
  SoftwareRendererEnabled = softwareRenderer;
}

void framebuffer_size_callback(GLFWwindow *window, int newWidth, int newHeight) {
  glViewport(0, 0, newWidth, newHeight);
  if (SoftwareRendererEnabled == true) {
    Vuelto::SoftwareRenderer::ResizeBuffer(newWidth, newHeight);
  }
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

  if (SoftwareRendererEnabled == true) {
    Vuelto::SoftwareRenderer::Init(window);
  }

  return window;
}

void DestroyWindow(Window win) { glfwDestroyWindow(win.window); }

void Terminate() { glfwTerminate(); }

}  // namespace Application

bool Window::WindowShouldClose() {
  while (!glfwWindowShouldClose(window)) {
    Window::WindowRefresh();
    if (Application::SoftwareRendererEnabled == true) Vuelto::SoftwareRenderer::Refresh();
    return false;
  }
  if (Application::SoftwareRendererEnabled == true) Vuelto::SoftwareRenderer::Terminate();
  glfwDestroyWindow(window);
  return true;
}

void Window::WindowRefresh() {
  glfwSwapBuffers(window);
  glfwPollEvents();
  glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
}
}  // namespace Vuelto
