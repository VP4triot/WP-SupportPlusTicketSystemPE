# WP-SupportPlusTicketSystemPE

This is a PoC written in Go for Wordpress Privilege Escalation. It was based on exploit´s logic of https://www.exploit-db.com/exploits/41006

Build it:  
go build PE-WPTicketSystem.go

Execute it:  
[Linux]  
PE-WPTicketSystem https://URL/wp-admin/admin-ajax.php user

[Windows]  
PE-WPTicketSystem.exe https://URL/wp-admin/admin-ajax.php user

You will need to add the cookies to your web request to get privileged access if vulnerable. Intended for educational purposes only.
