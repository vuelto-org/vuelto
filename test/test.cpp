#include "../src/Vuelto.hpp"

int main() {
  Vuelto::Init();

  Vuelto::Window win = Vuelto::Application::CreateWindow(800, 600, "test", false);

  Vuelto::SoftwareRenderer::Init(win);

  while (!win.WindowShouldClose()) {
    Vuelto::SoftwareRenderer::SetBackgroundColor(255, 255, 255);
    Vuelto::SoftwareRenderer::DrawRect(550, 550, 200, 200, 155, 155, 155);
    Vuelto::SoftwareRenderer::Refresh();
    Vuelto::Renderer2D::DrawRect(-0.1, -0.1, 0.5, 0.5, 0.1, 0.4, 0.7);
  }

  Vuelto::Application::DestroyWindow(win);
  Vuelto::SoftwareRenderer::Terminate();
  Vuelto::Terminate();

  return 0;
}
