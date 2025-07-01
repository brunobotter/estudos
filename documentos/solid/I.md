 ISP – Interface Segregation Principle

 Problema: Interfaces muito grandes que forçam implementações desnecessárias.

type Worker interface {
    Work()
    Eat()
}

type Human struct{}

func (h Human) Work() { fmt.Println("Working...") }
func (h Human) Eat()  { fmt.Println("Eating...") }

type Robot struct{}

func (r Robot) Work() { fmt.Println("Working...") }
func (r Robot) Eat()  { /* ??? Robots don't eat */ }


👉 Robot é forçado a implementar Eat, que não faz sentido.

type Workable interface {
    Work()
}

type Eatable interface {
    Eat()
}

type Human struct{}

func (h Human) Work() { fmt.Println("Working...") }
func (h Human) Eat()  { fmt.Println("Eating...") }

type Robot struct{}

func (r Robot) Work() { fmt.Println("Working...") }

✅ Agora:

Cada struct implementa apenas as interfaces que fazem sentido.