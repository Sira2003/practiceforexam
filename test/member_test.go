package test

import(
	"example.com/t1/entity"
	"testing"
	//"time"
	. `github.com/onsi/gomega`
	"github.com/asaskevich/govalidator"
	
)

func TestMember(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`UserName not null` ,func(t *testing.T) {
		member := entity.Member{
			UserName: "",
			Password: "123",
			Email: "kong@gmail.com",
		}
	ok, err := govalidator.ValidateStruct(member)
	
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("UserName not null"))
	})

	t.Run(`UserName is valid`, func(t *testing.T){
		member := entity.Member{
			UserName: "123",
			Password: "123",
			Email: "kong@gmail.com",
		}
		ok, err :=govalidator.ValidateStruct(member)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestEmail(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Email not null`,func(t *testing.T){
		member := entity.Member{
			UserName: "123",
			Password: "123",
			Email: "",
		}
		ok ,err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email not null"))
	})

	t.Run(`Email is invalid`, func(t *testing.T){
		member := entity.Member{
			UserName: "a",
			Password: "123",
			Email: "a@",
		}
		ok,err :=govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
	t.Run(`Email is correct`, func(t *testing.T){
		member := entity.Member{
			UserName: "123",
			Password: "123",
			Email: "kong@gmail.com",
		}
		ok,err :=govalidator.ValidateStruct(member)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}