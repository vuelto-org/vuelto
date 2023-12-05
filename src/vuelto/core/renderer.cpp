#include "renderer.hpp"

#include <iostream>
#include <vector>

#include "../tools/definitions.hpp"
#define STB_IMAGE_IMPLEMENTATION
#include "../vendor/stb_image/stb_image.h"
#include "app.hpp"

namespace Vuelto {
namespace Application {

Vuelto::Renderer2D CreateRenderer2D(Window win) {
  Vuelto::Renderer2D renderer;
  return renderer;
}

}  // namespace Application

Renderer2D::Image image_array[5];
int size = 0;

void Renderer2D::DrawRect(float x, float y, float height, float width, float color1, float color2, float color3) {
  glBegin(GL_QUADS);
  glColor3f(color1, color2, color3);
  glVertex2f(x, y);
  glVertex2f(x + width, y);
  glVertex2f(x + width, y + height);
  glVertex2f(x, y + height);
  glEnd();

  glColor3f(1.0f, 1.0f, 1.0f);
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

Renderer2D::Image Renderer2D::LoadImage(const char* imagePath, float x, float y, float width, float height) {
  Image image;

  int imgWidth, imgHeight, channels;
  unsigned char* imageData = stbi_load(imagePath, &imgWidth, &imgHeight, &channels, 4);  // Force 4 channels (RGBA)

  if (!imageData) {
    std::cerr << "Failed to load image: " << imagePath << std::endl;
    return image;
  }

  glGenTextures(1, &image.texture);
  glBindTexture(GL_TEXTURE_2D, image.texture);
  glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, imgWidth, imgHeight, 0, GL_RGBA, GL_UNSIGNED_BYTE, imageData);

  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);

  stbi_image_free(imageData);

  image.x = x;
  image.y = y;
  image.width = width;
  image.height = height;

  image_array[size++] = image;

  return image;
}

void Renderer2D::Image::DrawImage() {
  glBindTexture(GL_TEXTURE_2D, texture);

  glBegin(GL_QUADS);
  glTexCoord2f(0.0f, 0.0f);
  glVertex2f(x, y);
  glTexCoord2f(1.0f, 0.0f);
  glVertex2f(x + width, y);
  glTexCoord2f(1.0f, 1.0f);
  glVertex2f(x + width, y + height);
  glTexCoord2f(0.0f, 1.0f);
  glVertex2f(x, y + height);
  glEnd();

  glBindTexture(GL_TEXTURE_2D, 0);
}

void Renderer2D::CleanUp() {
  for (auto& image : image_array) {
    glDeleteTextures(1, &image.texture);
  }
}

}  // namespace Vuelto
