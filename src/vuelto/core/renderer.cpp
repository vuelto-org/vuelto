#include "renderer.hpp"

#include "../tools/definitions.hpp"
#include "app.hpp"

namespace Vuelto {

void Renderer2D::DrawRect(float x, float y, float height, float width, float color1, float color2, float color3) {
  glBegin(GL_QUADS);
  glColor3f(color1, color2, color3);
  glVertex2f(x, y);
  glVertex2f(x + width, y);
  glVertex2f(x + width, y + height);
  glVertex2f(x, y + height);
  glEnd();
}

void Renderer2D::SetBackgroundColor(float color1, float color2, float color3) {
  glClearColor(color1, color2, color3, 1);
}

void Renderer2D::DrawLine(int x1, int x2, int y1, int y2, float color1, float color2, float color3) {
  glLineWidth(1);
  glBegin(GL_LINES);
  glColor3f(color1, color2, color3);
  glVertex2f(x1, y1);
  glVertex2f(x2, y2);
  glEnd();
}

unsigned char* pixels;
int width;
int height;

void SoftwareRenderer::Init(int nwidth, int nheight) {
  width = nwidth * 2;
  height = nheight * 2;
  pixels = new unsigned char[width * height * 3];
}

void SoftwareRenderer::SetPixel(int x, int y, unsigned char r, unsigned char g, unsigned char b) {
  if (x >= 0 && x < width && y >= 0 && y < height) {
    int index = 3 * (y * width + x);
    pixels[index] = r;
    pixels[index + 1] = g;
    pixels[index + 2] = b;
  }
}

void SoftwareRenderer::DrawRect(int x, int y, int width, int height, int r, int g, int b) {
  for (int i = x; i < x + width; ++i) {
    for (int j = y; j < y + height; ++j) {
      SetPixel(i, j, r, g, b);
    }
  }
}

void SoftwareRenderer::SetBackgroundColor(unsigned char r, unsigned char g, unsigned char b) {
  for (int i = 0; i < width * height; ++i) {
    int index = 3 * i;
    pixels[index] = r;
    pixels[index + 1] = g;
    pixels[index + 2] = b;
  }
}

void SoftwareRenderer::Refresh() { glDrawPixels(width, height, GL_RGB, GL_UNSIGNED_BYTE, pixels); }

void SoftwareRenderer::Terminate() { delete[] pixels; }

void SoftwareRenderer::ResizeBuffer(int newWidth, int newHeight) {
  delete[] pixels;
  width = newWidth;
  height = newHeight;
  pixels = new unsigned char[width * height * 3];
}

}  // namespace Vuelto
