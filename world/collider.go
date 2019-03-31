package world

// import "github.com/harrisonthorne/chicky-chicky-go/maths"

// // Collider is a thing that can collide with other Colliders
// type Collider struct {
//     iCollider
//     Hitbox *maths.AABB
//     frozen bool // if true, collision fixes won't move the Collider
// }

// // CollidesWith returns whether or not the Collider
// // collides with another Collider
// func (c *Collider) CollidesWith(other *Collider) bool {
//     return c.Hitbox.CollidesWith(other.Hitbox)
// }

// // FixCollision fixes a collision between two
// // PhysicalObjects.
// func (c *Collider) FixCollision(other *Collider) {
//     if !c.CollidesWith(other) {
//         return
//     }

//     switch {
//     case !c.frozen && !other.frozen:

//     case !c.frozen && other.frozen:

//     case c.frozen && !other.frozen:

//     }

//     c.OnCollisionFix(*c.Hitbox)
// }

// type iCollider interface {
//     OnCollisionFix(newBounds maths.AABB)
// }
