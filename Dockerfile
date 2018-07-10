FROM scratch
ADD ba-pp-SecurePay ba-pp-SecurePay
EXPOSE 7777
ENTRYPOINT ["/ba-pp-SecurePay"]