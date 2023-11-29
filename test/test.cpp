#include "../src/Vuelto.hpp"

int main() {
  Vuelto::Application::Init();

  Vuelto::Window win1 = Vuelto::Application::CreateWindow(800, 600, "test", true);
  // Vuelto::Window win2 = Vuelto::Application::CreateWindow(800, 600, "test", true);

  while (!win1.WindowShouldClose()) {
    // win1.MakeContextCurrent();
    Vuelto::Renderer2D::DrawRect(-0.1, -0.1, 0.5, 0.5, 0.1, 0.4, 0.7);
    win1.Refresh();

    // win2.MakeContextCurrent();
    // Vuelto::Renderer2D::DrawRect(-0.1, -0.1, 0.5, 0.5, 0.1, 0.4, 0.7);
    // win2.WindowRefresh();
  }

  Vuelto::Application::Terminate();
  return 0;
}
