$('#cancel').click(function(){

  // redirect
  window.location.href = '/';

  // prevent default behavior of button
  return false;
});

/*
alertTimeout = setTimeout(function() {
	$('.alert').fadeOut("slow", "linear");
}, 5000);
/*
$('.alert').click(function(){
	$(this).fadeOut("slow", "linear");
	clearTimeout(alertTimeout);
});
*/
