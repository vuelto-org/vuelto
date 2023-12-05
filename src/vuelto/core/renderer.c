#include "renderer.h"

#include <stdio.h>

#include "../tools/definitions.h"
#define STB_IMAGE_IMPLEMENTATION
#include "../vendor/stb_image/stb_image.h"

Vuelto_Image image_array[5];
int size = 0;

void vueltoDrawRect(float x, float y, float height, float width, float color1, float color2, float color3) {
  glBegin(GL_QUADS);
  glColor3f(color1, color2, color3);
  glVertex2f(x, y);
  glVertex2f(x + width, y);
  glVertex2f(x + width, y + height);
  glVertex2f(x, y + height);
  glEnd();

  glColor3f(1.0f, 1.0f, 1.0f);
}

void vueltoSetBackgroundColor(float color1, float color2, float color3) { glClearColor(color1, color2, color3, 1); }

void vueltoDrawLine(int x1, int x2, int y1, int y2, float color1, float color2, float color3) {
  glLineWidth(1);
  glBegin(GL_LINES);
  glColor3f(color1, color2, color3);
  glVertex2f(x1, y1);
  glVertex2f(x2, y2);
  glEnd();
}

Vuelto_Image vueltoLoadImage(const char* imagePath, float x, float y, float width, float height) {
  Vuelto_Image image;

  int imgWidth, imgHeight, channels;
  unsigned char* imageData = stbi_load(imagePath, &imgWidth, &imgHeight, &channels, 4);  // Force 4 channels (RGBA)

  if (!imageData) {
    printf("Failed to load image: %s \n", imagePath);
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

void vueltoDrawImage(Vuelto_Image img) {
  glBindTexture(GL_TEXTURE_2D, img.texture);

  glBegin(GL_QUADS);
  glTexCoord2f(0.0f, 0.0f);
  glVertex2f(img.x, img.y);
  glTexCoord2f(1.0f, 0.0f);
  glVertex2f(img.x + img.width, img.y);
  glTexCoord2f(1.0f, 1.0f);
  glVertex2f(img.x + img.width, img.y + img.height);
  glTexCoord2f(0.0f, 1.0f);
  glVertex2f(img.x, img.y + img.height);
  glEnd();

  glBindTexture(GL_TEXTURE_2D, 0);
}

void vueltoCleanUp() {
  vueltoForEach(Vuelto_Image * i, image_array) { glDeleteTextures(1, &i->texture); }
}
