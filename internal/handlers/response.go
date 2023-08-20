package handlers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

func MapToStruct(c *fiber.Ctx, input interface{}) error {
	if input == nil {
		fmt.Println("ooooo99")
		// ทำสิ่งที่ต้องการเมื่อ input เป็น nil
	}
	contentType := c.Get("Content-Type")
	switch {
	case strings.HasPrefix(contentType, "application/json"):
		if err := c.BodyParser(input); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"), strings.HasPrefix(contentType, "multipart/form-data"):
		if err := mapFormValues(c, input); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		if err := c.QueryParser(input); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported content type: %s", contentType)
	}

	// Validate ค่าที่รับเข้ามา
	if err := validate.Struct(input); err != nil {
		// กรณีมี error ในการ validate แสดงข้อความเพิ่มเติม
		errs := err.(validator.ValidationErrors)
		errorMsg := "Invalid request data:"
		for _, e := range errs {
			errorMsg += fmt.Sprintf("\n- Field: %s, Type: %T, Error: %s", e.Field(), e.Value(), e.Tag())
		}
		return fmt.Errorf(errorMsg)
	}

	return nil
}

func ResponseObject(c *fiber.Ctx, fn interface{}, input interface{}) error {
	contentType := c.Get("Content-Type")
	switch {
	case strings.HasPrefix(contentType, "application/json"):
		if err := c.BodyParser(input); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"), strings.HasPrefix(contentType, "multipart/form-data"):
		if err := mapFormValues(c, input); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		if err := c.QueryParser(input); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported content type: %s", contentType)
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
		reflect.ValueOf(input),
	})

	errObj := out[1].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return errObj.(error)
	}

	return RenderJSON(c, out[0].Interface())
}

func ResponseObjectNoRequest(c *fiber.Ctx, fn interface{}) error {
	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
	})
	errObj := out[1].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return errObj.(error)
	}

	return RenderJSON(c, out[0].Interface())
}

func MapToStruct_xw(c *fiber.Ctx, input interface{}) error {
	contentType := c.Get("Content-Type")
	switch {
	case strings.HasPrefix(contentType, "application/json"):
		if err := c.BodyParser(input); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"), strings.HasPrefix(contentType, "multipart/form-data"):
		if err := mapFormValues(c, input); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		if err := c.QueryParser(input); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported content type: %s", contentType)
	}

	// Validate ค่าที่รับเข้ามา
	if err := validate.Struct(input); err != nil {
		// กรณีมี error ในการ validate แสดงข้อความเพิ่มเติม
		errs := err.(validator.ValidationErrors)
		errorMsg := "Invalid request data:"
		for _, e := range errs {
			errorMsg += fmt.Sprintf("\n- Field: %s, Type: %T, Error: %s", e.Field(), e.Value(), e.Tag())
		}
		return fmt.Errorf(errorMsg)
	}

	return nil
}

func MapToStructX(c *fiber.Ctx, fn func(*fiber.Ctx, interface{}) error, input interface{}) error {
	contentType := c.Get("Content-Type")
	switch {
	case strings.HasPrefix(contentType, "application/json"):
		if err := c.BodyParser(input); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"), strings.HasPrefix(contentType, "multipart/form-data"):
		if err := mapFormValues(c, input); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		if err := c.QueryParser(input); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported content type: %s", contentType)
	}

	// Validate ค่าที่รับเข้ามา
	if err := validate.Struct(input); err != nil {
		// กรณีมี error ในการ validate แสดงข้อความเพิ่มเติม
		errs := err.(validator.ValidationErrors)
		errorMsg := "Invalid request data:"
		for _, e := range errs {
			errorMsg += fmt.Sprintf("\n- Field: %s, Type: %T, Error: %s", e.Field(), e.Value(), e.Tag())
		}
		return fmt.Errorf(errorMsg)
	}

	// // Call the func with the mapped input
	// out := reflect.ValueOf(fn).Call([]reflect.Value{
	// 	reflect.ValueOf(c),
	// 	reflect.ValueOf(input),
	// })

	// errObj := out[1].Interface()
	// if errObj != nil {
	// 	logrus.Errorf("call service error: %s", errObj)
	// 	return errObj.(error)
	// }

	// return RenderJSON(c, out[0].Interface())
	return fn(c, input)
}

func mapFormValues(c *fiber.Ctx, input interface{}) error {
	val := reflect.ValueOf(input).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		value := c.FormValue(fieldName)
		if value != "" {
			// If the field is a slice, assign the value as a single element
			if field.Type().Kind() == reflect.Slice {
				slice := reflect.MakeSlice(field.Type(), 1, 1)
				slice.Index(0).SetString(value)
				field.Set(slice)
			} else {
				field.SetString(value)
			}
		}
	}

	return nil
}

func RenderJSON(c *fiber.Ctx, response interface{}) error {
	return c.
		Status(fiber.StatusOK).
		JSON(response)
}
