#include <GLFW/glfw3.h>

#include "OpenGL/gl.h"

unsigned char* pixels;
int width;
int height;

void Init(int initialWidth, int initialHeight) {
  width = initialWidth * 2;
  height = initialHeight * 2;
  pixels = new unsigned char[width * height * 3];
}

void SetPixel(int x, int y, unsigned char r, unsigned char g, unsigned char b) {
  if (x >= 0 && x < width && y >= 0 && y < height) {
    int index = 3 * (y * width + x);
    pixels[index] = r;
    pixels[index + 1] = g;
    pixels[index + 2] = b;
  }
}

void DrawRect(int x, int y, int width, int height, int r, int g, int b) {
  for (int i = x; i < x + width; ++i) {
    for (int j = y; j < y + height; ++j) {
      SetPixel(i, j, r, g, b);
    }
  }
}

void SetBackgroundColor(unsigned char r, unsigned char g, unsigned char b) {
  for (int i = 0; i < width * height; ++i) {
    int index = 3 * i;
    pixels[index] = r;
    pixels[index + 1] = g;
    pixels[index + 2] = b;
  }
}

void Refresh() { glDrawPixels(width, height, GL_RGB, GL_UNSIGNED_BYTE, pixels); }

void Terminate() { delete[] pixels; }

int main() {
  // Initialize GLFW
  if (!glfwInit()) {
    return -1;
  }

  // Create a windowed mode window and its OpenGL context
  GLFWwindow* window = glfwCreateWindow(800, 600, "Simple Renderer", NULL, NULL);
  if (!window) {
    glfwTerminate();
    return -1;
  }

  // Make the window's context current
  glfwMakeContextCurrent(window);

  // Initialize SoftwareRenderer
  Init(800, 600);

  // Loop until the user closes the window
  while (!glfwWindowShouldClose(window)) {
    // Set the background color
    SetBackgroundColor(255, 0, 0);  // Red background

    // Draw a white rectangle
    DrawRect(100, 100, 200, 150, 255, 255, 255);  // White rectangle

    // Refresh the display
    Refresh();

    // Swap front and back buffers
    glfwSwapBuffers(window);

    // Poll for and process events
    glfwPollEvents();
  }

  // Terminate SoftwareRenderer
  Terminate();

  // Terminate GLFW
  glfwTerminate();

  return 0;
}
