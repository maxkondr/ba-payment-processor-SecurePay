FROM scratch
ADD ba-payment-processor-A ba-payment-processor-A
EXPOSE 7777
ENTRYPOINT ["/ba-payment-processor-A"]