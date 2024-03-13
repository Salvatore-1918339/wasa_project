<script>
export default {
	props: ['msg']
}
</script>

<!--Questo codice visualizza un messaggio di errore in una finestra di dialogo di avviso. 
    Il codice HTML utilizza la direttiva Vue v-if per visualizzare il messaggio di errore solo quando esiste un messaggio
    da visualizzare. Il messaggio viene visualizzato all'interno di un elemento div con la classe "alert alert-danger" per indicare
    che si tratta di un messaggio di errore.

Il codice CSS definisce la posizione del messaggio di errore, stabilendo che deve essere posizionato assolutamente
nella parte inferiore della pagina con una larghezza del 50% centrato rispetto alla pagina.-->

<template>
	<div class="d-flex justify-content-center alert alert-danger mt-3 error" role="alert">
		{{ msg }}
	</div>
</template>

<style>
.error {
    position: absolute;
    width: 50%;
    left: 25%;
    bottom: 0%;
}
</style>