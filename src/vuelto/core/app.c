#include "app.h"

#include "../tools/definitions.h"
#include "renderer.h"

bool MultipleWindowsEnabled = false;

void vueltoInit() {
  if (!glfwInit()) {
    printf("GLFW Init failed\n");
  }
}

void vueltoInitMultipleWindows() {
  vueltoInit();
  MultipleWindowsEnabled = true;
}

void framebuffer_size_callback(GLFWwindow *window, int newWidth, int newHeight) {
  glViewport(0, 0, newWidth, newHeight);
}

Vuelto_Window vueltoCreateWindow(int width, int height, const char *title, bool resizable) {
  glfwWindowHint(GLFW_RESIZABLE, resizable);

  GLFWwindow *glfw_window = glfwCreateWindow(width, height, title, NULL, NULL);
  if (!glfw_window) {
    glfwTerminate();
    printf("GLFW Window creation failed\n");
  }

  glfwSetFramebufferSizeCallback(glfw_window, framebuffer_size_callback);

  if (!MultipleWindowsEnabled) glfwMakeContextCurrent(glfw_window);

  glEnable(GL_BLEND);
  glEnable(GL_TEXTURE_2D);

  glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA);

  Vuelto_Window window;
  window.width = width;
  window.height = height;
  window.title = title;
  window.window = glfw_window;

  return window;
}

void vueltoDestroyWindow(Vuelto_Window win) { glfwDestroyWindow(win.window); }

void vueltoTerminate() { glfwTerminate(); }

bool vueltoWindowShouldClose(Vuelto_Window win) {
  while (!glfwWindowShouldClose(win.window)) {
    return false;
  }
  vueltoCleanUp();
  glfwDestroyWindow(win.window);
  return true;
}

void vueltoMakeContextCurrent(Vuelto_Window win) { glfwMakeContextCurrent(win.window); }

void vueltoRefresh(Vuelto_Window win) {
  glfwSwapBuffers(win.window);
  if (!MultipleWindowsEnabled) glfwPollEvents();
  glClear(GL_COLOR_BUFFER_BIT);
}
