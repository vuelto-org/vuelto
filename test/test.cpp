#include "../src/Vuelto.hpp"

int main() {
  Vuelto::Application::Init();

  Vuelto::Window win = Vuelto::Application::CreateWindow(800, 600, "test", true);

  Vuelto::Renderer2D renderer = Vuelto::Application::CreateRenderer2D(win);

  Vuelto::Renderer2D::Image img = renderer.LoadImage("test/image.png", 100.0f, 100.0f, 200.0f, 150.0f);

  while (!win.WindowShouldClose()) {
    renderer.DrawRect(0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1);
    img.DrawImage();
    win.Refresh();
  }

  Vuelto::Application::Terminate();
  return 0;
}
