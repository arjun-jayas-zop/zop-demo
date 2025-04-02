package template

const Header = `
<header style="text-align: right; padding: 10px;">
    <img src="https://raw.githubusercontent.com/aryanmehrotra/gofr0-example/refs/heads/main/svgviewer-png-output.png" 
         alt="Logo" 
         style="height: 50px; width: auto; display: block; margin-left: auto; margin-right: 0;">
</header>

<style>
    @media (max-width: 600px) {
        header {
            text-align: center !important;
        }
        header img {
            margin-left: auto !important;
            margin-right: auto !important;
            display: block !important;
        }
    }
</style>
`

const Footer = `<footer style="text-align: center; font-size: 0.9em; color: #aaaaaa; padding: 20px;">
    © 2024 ZopDev Technology Pvt. Ltd. All rights reserved.
</footer>

`

const WelcomeEmailURL = `https://zop.dev/app/account-verify?token=%v`
const WelcomeEmailSubject = `Welcome to Zop.dev!`
const WelcomeEmail = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Welcome to Zop.dev</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
<p>Hi %v,</p>
<p>Welcome to Zop.dev! We're excited to have you join our community of developers and teams simplifying cloud infrastructure management.</p>
<p>Your account has been successfully created. Here's what you can expect from Zop.dev:</p>
<ul>
<li>Effortless cloud management, empowering your developers.</li>
<li>Full control over your cloud environment through simple workflows.</li>
<li>Robust security and automation, all in one unified platform.</li>
</ul>
<p>To get started, click the button below to verify your email and get started to set up your cloud infrastructure:</p>
<a href="%v" style="display: inline-block; padding: 10px 20px; color: #fff; background-color: #0A9BBB; text-decoration: none; border-radius: 5px;">Verify Email</a>
<p>If you have any questions or need assistance, please refer to our <a href="https://zop.dev/documentation" style="color: #0A9BBB;">documentation</a> or <a href="https://zop.dev/support" style="color: #0A9BBB;">contact our support team</a>.</p>
<p>Welcome aboard,</p>
<p>The Zop Team</p>
</body>
</html>
`

const PasswordResetEmailURL = `https://zop.dev/app/password-reset?token=%v`
const PasswordResetEmailSubject = `Reset Your Zop Password`
const PasswordResetEmail = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Password Reset</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <p>Hi %v,</p>
    <p>You recently requested to reset your Zop account password.</p>
    <p>To proceed with the password reset, click the button below:</p>
	<a href="%v" style="display: inline-block; padding: 10px 20px; color: #fff; background-color: #0A9BBB; text-decoration: none; border-radius: 5px;">Reset Password</a>
    <p>This password reset link is valid for the next 15 minutes. If you do not reset your password within this time, you will need to submit a new request.</p>
    <p>If you did not request a password reset, please ignore this email. Your password will remain unchanged.</p>
    <p>For any questions or concerns, please <a href="https://zop.dev/support" style="color: #0A9BBB; text-decoration: none;">contact our support team</a>.</p>
    <p>Best,</p>
    <p>The Zop Team</p>
</body>
</html>
`

const PasswordVerificationResendEmailURL = WelcomeEmailURL
const PasswordVerificationResendEmailSubject = `Verify Your Zop Email`
const PasswordVerificationResendEmail = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Resend Password Verification Link</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <p>Hi %v,</p>
    <p>You recently requested to resend your email verification link.</p>
    <p>To verify your email and activate your account, click the button below:</p>
    <a href="%v" style="display: inline-block; padding: 10px 20px; color: #fff; background-color: #0A9BBB; text-decoration: none; border-radius: 5px;">Verify Email</a>
    <p>This verification link is valid for the next 15 minutes. If you do not complete verification within this time, you will need to request a new link.</p>
    <p>If you did not request this email, please ignore it. Your account and password will remain unchanged.</p>
    <p>For any questions or assistance, please <a href="https://zop.dev/support" style="color: #0A9BBB;">contact our support team</a>.</p>
    <p>Thank you,</p>
    <p>The Zop Team</p>
</body>
</html>
`

const SecurityAlertEmailSubject = `Security Alert: Password Change Notification`
const SecurityAlertEmail = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Password Change Notification</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <p>Hi %v,</p>
    <p>We wanted to let you know that your Zop account password was successfully changed.</p>
    <p>If you made this change, you can safely ignore this message. However, if you did not initiate this password change, please revert to this email immediately.</p>
    
    <p>Thank you,</p>
    <p>The Zop Team</p>
</body>
</html>
`
