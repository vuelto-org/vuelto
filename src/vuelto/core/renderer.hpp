#pragma once
#include "../tools/definitions.h"
#include "app.hpp"

namespace Vuelto {
namespace Renderer2D {

void DrawRect(float x, float y, float height, float width, float color1, float color2, float color3);
void SetBackgroundColor(float color1, float color2, float color3);
void DrawLine(int x1, int x2, int y1, int y2, float color1, float color2, float color3);

}  // namespace Renderer2D

namespace SoftwareRenderer {

void Init(Vuelto::Window window);
void SetPixel(int x, int y, unsigned char r, unsigned char g, unsigned char b);
void DrawRect(int x, int y, int width, int height, int r, int g, int b);
void SetBackgroundColor(unsigned char r, unsigned char g, unsigned char b);
void Refresh();
void Terminate();
void ResizeBuffer(int newWidth, int newHeight);
bool Initialized();
}  // namespace SoftwareRenderer

}  // namespace Vuelto
