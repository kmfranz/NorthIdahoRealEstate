$(document).ready(function(){


	$('.photo-grid-container div').hover(function(){
		$(this).find('.hover-box').fadeIn(100)
		$(this).find('.photo-subtitle').fadeOut(100)
	}, function(){
		$(this).find('.hover-box').fadeOut(100)
		$(this).find('.photo-subtitle').fadeIn(100)
	})

	$('.kf-navbar span img').hover(function(){
		$(this).attr('src', '/resources/images/mini logo hover.png');
	}, function(){
		$(this).attr('src', '/resources/images/mini logo.png');
	})

	$('.city-images-mobile div').hover(function(){
		$(this).find('.hover-div').fadeIn(100)
		$(this).find('.hover-text').css('color', 'black')
	}, function(){
		$(this).find('.hover-div').fadeOut(100)
		$(this).find('.hover-text').css('color', 'white')
	})

	//about page
	$('.nav-bar-white span img').hover(function(){
		$(this).attr('src', '/resources/images/mini logo.png');
	}, function(){
		$(this).attr('src', '/resources/images/mini logo hover.png');
	})


	//resize small cells
	//console.log($(this).width())

});


//resizing the custom image cells

$(window).resize(function(){
	//console.log($(this).width())
})
