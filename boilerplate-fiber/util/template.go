package util

import (
	"fmt"
)

type templateStruct struct{}

func NewTemplate() *templateStruct {
	return &templateStruct{}
}

func (t *templateStruct) EmailResetPassword(link string, expired string) (template string) {
	template = fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		
		<head>
			<meta charset="UTF-8" />
			<meta http-equiv="X-UA-Compatible" content="IE=edge" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<link rel="stylesheet"
				href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap" />
		
			<title>Akun anda telah dibuat</title>
		</head>
		<style>
			body {
				font-family: "Poppins", sans-serif;
				/* height: 100vh; */
			}
		</style>
		
		<body style="
			text-align: center;
			color: #1f2f59;
			width: 600px;
			margin: auto;
			word-wrap: break-word;
			background-color: whitesmoke;
			">
			<div class="container"
				style="width: 600px; border: 1px solid lightgrey;  background-color: white; padding-top: 120px;">
				<div style="padding: 50px 20px 150px 20px;">
					<h2>
						SISAMBI
					</h2>
					<hr style="
					height: 6px;
					width: 70px;
					border: none;
					background-color: #047CD1;
					border-radius: 5px;
					margin: auto;
				" />
					<div>
						<p>Berikut adalah link untuk melakukan reset password: <a href="%v">%v</a></p>
						<p>hanya berlaku sampai %v, dan hanya bisa digunakan satu kali.</p>
					</div>
					
				
				</div>
				<footer style="
				background-color: #047CD1;
				color: white;
				font-weight: 400;
				padding: 10px 0px;
				margin-top: 30px;
				bottom: 0;
				
				">
					<p>Email sistem SISAMBI</p>
					<p>Mohon jangan dibalas</p>
				</footer>
			</div>
		</body>
		
		</html>
	`, link, link, expired)
	return
}
