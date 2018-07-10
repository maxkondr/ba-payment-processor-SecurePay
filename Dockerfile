FROM scratch
ADD ba-payment-processor-secure-pay ba-payment-processor-secure-pay
EXPOSE 7777
ENTRYPOINT ["/ba-payment-processor-secure-pay"]