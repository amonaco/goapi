package notification

const userCreateTemplate = `
<p>Dear {{.Name}}, please click in the following link in order to
	<a href="{{.TokenURL}}">
		setup your account.
	</a>
</p>`

const forgotPasswordTemplate = `
<p>Dear {{.Name}}, please click in the following link in order to
	<a href="{{.TokenURL}}">
		reset your password.
	</a>
</p>`
