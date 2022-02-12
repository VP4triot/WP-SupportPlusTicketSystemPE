# WP-SupportPlusTicketSystemPE

This is a PoC written in Go for Wordpress Privilege Escalation. It was based on exploit´s logic of https://www.exploit-db.com/exploits/41006

Build it:
go build PE-Wordpress.go

Execute it:
[Linux] 
PE-Wordpress https://URL/wp-admin/admin-ajax.php user

[Windows] 
PE-Wordpress.exe https://URL/wp-admin/admin-ajax.php user
