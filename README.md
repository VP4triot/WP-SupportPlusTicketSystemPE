# WP-SupportPlusTicketSystemPE

This is a PoC written in Go for Wordpress Privilege Escalation. It was based on exploit´s logic of https://www.exploit-db.com/exploits/41006

Build it:
go build PE-WPTicketSystem.go

Execute it:
[Linux] 
PE-WPTicketSystem.go https://URL/wp-admin/admin-ajax.php user

[Windows] 
PE-WPTicketSystem.go https://URL/wp-admin/admin-ajax.php user

This is a PoC you will need to add the cookies to your web request to get privileged access if vulnerable.
