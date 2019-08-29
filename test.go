package main

import (
    "fmt"
    "github.com/go-playground/locales/zh"
    ut "github.com/go-playground/universal-translator"
    "gopkg.in/go-playground/validator.v9"
    zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

// User contains user information
type User struct {
    FirstName      string     `validate:"required"`
    LastName       string     `validate:"required"`
    Age            uint8      `validate:"gte=0,lte=130"`
    Email          string     `validate:"required,email"`
    FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
    Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
    Street string `validate:"required"`
    City   string `validate:"required"`
    Planet string `validate:"required"`
    Phone  string `validate:"required"`
}

// use a single instance , it caches struct info
var (
    uni      *ut.UniversalTranslator
    validate *validator.Validate
)

func main() {
    
    // NOTE: ommitting allot of error checking for brevity

    zh := zh.New()
    uni = ut.New(zh)
    
    // this is usually know or extracted from http 'Accept-Language' header
    // also see uni.FindTranslator(...)
    trans, _ := uni.GetTranslator("zh")
    
    validate = validator.New()
    zh_translations.RegisterDefaultTranslations(validate, trans)
    
    translateAll(trans)
    translateIndividual(trans)
}

func translateAll(trans ut.Translator) {
    
    type User struct {
        Username string `validate:"required"`
        Tagline  string `validate:"required,lt=10"`
        Tagline2 string `validate:"required,gt=1"`
    }
    
    user := User{
        Username: "Joeybloggs",
        Tagline:  "This tagline is way too long.",
        Tagline2: "",
    }
    
    err := validate.Struct(user)
    if err != nil {
        
        // translate all error at once
        errs := err.(validator.ValidationErrors)
        
        // returns a map with key = namespace & value = translated error
        // NOTICE: 2 errors are returned and you'll see something surprising
        // translations are i18n aware!!!!
        // eg. '10 characters' vs '1 character'
        fmt.Println(errs.Translate(trans))
    }
}

func translateIndividual(trans ut.Translator) {
    
    type User struct {
        Username string `validate:"required"`
    }
    
    var user User
    
    err := validate.Struct(user)
    if err != nil {
        
        errs := err.(validator.ValidationErrors)
        
        for _, e := range errs {
            // can translate each error one at a time.
            fmt.Println(e.Translate(trans))
        }
    }
}