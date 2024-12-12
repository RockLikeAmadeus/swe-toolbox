Sources: 
- https://postmarkapp.com/guides/everything-you-need-to-know-about-smtp
- https://mailtrap.io/blog/smtp/
- https://notes.eatonphil.com/handling-email-from-gmail-smtp-protocol-basics.html

An SMTP server handles the sending, receiving, and relaying of email. 

Look into https://www.loginradius.com/blog/engineering/sending-emails-with-golang/
https://api-docs.mailtrap.io/docs/mailtrap-api-docs/127a373cd0493-receive-events
https://www.reddit.com/r/softwaretesting/comments/1d4z2yb/what_do_you_use_for_automating_email_testing/

https://greenmail-mail-test.github.io/greenmail/ <-- with Docker?

At the end of the day, ended up using a project called mailhog, which worked very well, and has a ready made Docker image