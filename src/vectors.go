package vuelto

type Vector2D struct {
	X float64
	Y float64
}

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (a *Application) NewVector2D(x, y float64) Vector2D {
	return Vector2D{x, y}
}

func (a *Application) AddVector2D(v1, v2 Vector2D) Vector2D {
	return Vector2D{v1.X + v2.X, v1.Y + v2.Y}
}

func (a *Application) SubtractVector2D(v1, v2 Vector2D) Vector2D {
	return Vector2D{v1.X - v2.X, v1.Y - v2.Y}
}

func (a *Application) NewVector3D(x, y, z float64) Vector3D {
	return Vector3D{x, y, z}
}

func (a *Application) AddVector3D(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (a *Application) SubtractVector3D(v1, v2 Vector3D) Vector3D {
	return Vector3D{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}
