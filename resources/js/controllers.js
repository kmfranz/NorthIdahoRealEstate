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

	$('.photo-grid-container .photo-link').click(function(){

		link_id = $(this).attr('id')
		cur_href = window.location.href;

		window.location.href = cur_href + link_id;

	});

	$('.city-images-mobile .mobile-city-link').click(function(){

		link_id = $(this).attr('id')
		cur_href = window.location.href;

		window.location.href = cur_href + link_id;

	});


	$('.mobile-current-city').click(function(){
		console.log('Here')
		if(!$('.rotate').hasClass('up')){
			$('.hidden-city-menu').slideDown({duration: 600, easing:"easeOutQuart"})
			$('.rotate').toggleClass("up");
		} else{
			$('.hidden-city-menu').slideUp({duration: 600, easing:"easeOutQuart"})
			$('.rotate').toggleClass("up");
		}
	})


	//resize small cells
	//console.log($(this).width())

});

function loadPage(page_name){
	window.location.href = window.location.protocol + "//" + window.location.host + "/" + page_name;
}

//resizing the custom image cells

$(window).resize(function(){
	//console.log($(this).width())
})
