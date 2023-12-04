#include "../src/Vuelto.hpp"

int main() {
  Vuelto::Application::Init();

  Vuelto::Window win = Vuelto::Application::CreateWindow(800, 600, "test", true);

  Vuelto::Renderer2D renderer = Vuelto::Application::CreateRenderer2D(win);

  // Vuelto::Renderer2D::Image img = renderer.LoadImage("test/coin.png", -0.2, -0.2, 0.7, 0.7);

  while (!win.WindowShouldClose()) {
    renderer.DrawRect(0, 0, 0.7, 0.7, 0.5, 0.1, 0.3);
    // img.DrawImage();
    win.Refresh();
  }

  renderer.CleanUp();

  Vuelto::Application::Terminate();
  return 0;
}
