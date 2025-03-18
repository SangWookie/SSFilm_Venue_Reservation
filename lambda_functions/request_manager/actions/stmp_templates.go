package actions

import (
	"bytes"
	"text/template"
)

const reservationCompleteTemplate = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>예약 확인</title>
<style>
body {
    font-family: Arial, sans-serif;
    background-color: #f4f4f4;
    margin: 0;
    padding: 20px;
}
.container {
    max-width: 600px;
    background: #ffffff;
    padding: 25px;
    border-radius: 10px;
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
    margin: auto;
}
h2 {
    color: #007bff;
    text-align: center;
    margin-bottom: 20px;
}
.details {
    padding: 20px;
    background: #f8f9fa;
    border-left: 5px solid #007bff;
    border-radius: 8px;
    line-height: 1.7;
}
.details p {
    margin: 12px 0;
    font-size: 17px;
    color: #333;
}
.label {
    font-weight: bold;
    color: #007bff;
}
</style>
</head>
<body>
<div class="container">
    <h2>예약이 확정되었습니다.</h2>
    <div class="details">
        <p><span class="label">이름:</span> {{.Name}}</p>
        <p><span class="label">장소:</span> {{.Location}}</p>
        <p><span class="label">시간:</span> {{.Time}}</p>
        <p><span class="label">카테고리:</span> {{.Category}}</p>
        <p><span class="label">세부사항:</span> {{.Details}}</p>
    </div>
</div>
</body>
</html>
`
const reservationModifiedTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>예약 변경 안내</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 20px;
        }
        .container {
            max-width: 600px;
            background: #ffffff;
            padding: 25px;
            border-radius: 10px;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
            margin: auto;
        }
        h2 {
            color: #6f42c1;
            text-align: center;
            margin-bottom: 20px;
        }
        .details {
            padding: 20px;
            background: #f8f9fa;
            border-left: 5px solid #6f42c1;
            border-radius: 8px;
            line-height: 1.7;
        }
        .details p {
            margin: 12px 0;
            font-size: 17px;
            color: #333;
        }
        .label {
            font-weight: bold;
            color: #6f42c1;
        }
        .changed {
            font-weight: bold;
            color: #6f42c1;
        }
        .old {
            text-decoration: line-through;
            color: #888;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>예약이 변경되었습니다.</h2>
        <div class="details">
            <p><span class="label">이름:</span> {{.Name}}</p>
			<p><span class="label">장소:</span> {{.Location}}</p>
            <p><span class="label">시간:</span><span class="changed"> {{.Time}}</span></p>
            <p><span class="label">카테고리:</span> {{.Category}}</p>
            <p><span class="label">세부사항:</span> {{.Details}}</p>
        </div>
    </div>
</body>
</html>
`
const reservationCanceledTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>예약 취소 안내</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 20px;
        }
        .container {
            max-width: 600px;
            background: #ffffff;
            padding: 25px;
            border-radius: 10px;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
            margin: auto;
        }
        h2 {
            color: #dc3545;
            text-align: center;
            margin-bottom: 20px;
        }
        .details {
            padding: 20px;
            background: #f8f9fa;
            border-left: 5px solid #dc3545;
            border-radius: 8px;
            line-height: 1.7;
        }
        .details p {
            margin: 12px 0;
            font-size: 17px;
            color: #333;
        }
        .label {
            font-weight: bold;
            color: #dc3545;
        }
        .changed {
            font-weight: bold;
            color: #28a745;
        }
        .old {
            text-decoration: line-through;
            color: #888;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>예약이 취소되었습니다.</h2>
        <div class="details">
            <p><span class="label">이름:</span> {{.Name}}</p>
			<p><span class="label">장소:</span> {{.Location}}</p>
            <p><span class="label">시간:</span><span class="changed">{{.Time}}</span></p>
            <p><span class="label">카테고리:</span> {{.Category}}</p>
            <p><span class="label">세부사항:</span> {{.Details}}</p>
        </div>
    </div>
</body>
</html>
`

type ReservationEmailData struct {
	Name     string
	Location string
	Time     string
	Category string
	Details  string
}

func getReservationEmail(name, location, time, category, details string, style string) (string, error) {
	tmpl, err := template.New("reservationEmail").Parse(style)
	if err != nil {
		return "", err
	}

	data := ReservationEmailData{
		Name:     name,
		Location: location,
		Time:     time,
		Category: category,
		Details:  details,
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, data)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func GetReservationCompleteTemplate(data ReservationEmailData) (string, error) {
	return getReservationEmail(data.Name, data.Location, data.Time, data.Category, data.Details, reservationCompleteTemplate)
}
func GetReservationModifiedTemplate(data ReservationEmailData) (string, error) {
	return getReservationEmail(data.Name, data.Location, data.Time, data.Category, data.Details, reservationModifiedTemplate)
}
func GetReservationCanceledTemplate(data ReservationEmailData) (string, error) {
	return getReservationEmail(data.Name, data.Location, data.Time, data.Category, data.Details, reservationCanceledTemplate)
}
