document.addEventListener('DOMContentLoaded',async() =>{
    console.log("DOM fully loaded");
    const {publishableKey} = await fetch("https://ev-platform-server-production.up.railway.app/payment/config").then((r) => r.json());
    const stripe = Stripe(publishableKey);

    const {clientSecret} = await fetch("https://ev-platform-server-production.up.railway.app/payment/create-payment-intent",{
        method: "POST",
        headers:{
            "Content-Type":"application/json"
        },
        body: JSON.stringify({
            cost : 5000,
        }),
    }).then(r => r.json())

    const elements = stripe.elements({clientSecret})
    const paymentElement = elements.create('payment')
    paymentElement.mount('#payment-element')
    const form =document.getElementById('payment-form')
    form.addEventListener('submit',async (e) => {
        e.preventDefault();

        const {error} = await stripe.confirmPayment({
            elements,
            confirmParams:{
                return_url:  "https://ev-charging-frontend-cza3hgccgmb9f7bh.canadacentral-01.azurewebsites.net/complete.html"
            }
        })
        if(error){
            const messages = document.getElementById('error-messages')
            messages.innerText = error.message; 
        }
    })
});
function GetURLParameter(sParam)
{
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
    for (var i = 0; i < sURLVariables.length; i++) 
    {
        var sParameterName = sURLVariables[i].split('=');
        if (sParameterName[0] == sParam) 
        {
            return sParameterName[1];
        }
    }
}​
