$('#cancel').click(function(){

  // redirect
  window.location.href = '/';

  // prevent default behavior of button
  return false;
});


alertTimeout = setTimeout(function() {
	$('.alert-success').fadeOut("slow", "linear");
}, 3000);
/*
$('.alert').click(function(){
	$(this).fadeOut("slow", "linear");
	clearTimeout(alertTimeout);
});
*/
