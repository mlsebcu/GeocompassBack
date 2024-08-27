package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username"`
	Email    string `gorm:"not null; unique_index" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

// BeforeSave es un hook que se ejecuta antes de guardar el usuario en la base de datos
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	log.Println("Contraseña encriptada antes de guardar:", u.Password)
	return nil
}

// ComparePassword compara una contraseña en texto plano con la contraseña encriptada
func (u *User) ComparePassword(password string) bool {
	log.Println("Contraseña en la base de datos:", u.Password)
	log.Println("Contraseña ingresada para comparar:", password)

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Println("Error al comparar la contraseña:", err)
		return false
	}
	return true
}
