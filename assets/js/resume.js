(function($) {
    "use strict"; // Start of use strict

    // Smooth scrolling using jQuery easing
    $('a.js-scroll-trigger[href*="#"]:not([href="#"])').click(function() {
        if (location.pathname.replace(/^\//, '') == this.pathname.replace(/^\//, '') && location.hostname == this.hostname) {
            var target = $(this.hash);
            target = target.length ? target : $('[name=' + this.hash.slice(1) + ']');
            if (target.length) {
                $('html, body').animate({
                    scrollTop: (target.offset().top)
                }, 1000, "easeInOutExpo");
                return false;
            }
        }
    });

    // Closes responsive menu when a scroll trigger link is clicked
    $('.js-scroll-trigger').click(function() {
        $('.navbar-collapse').collapse('hide');
    });

    // Activate scrollspy to add active class to navbar items on scroll
    //$('body').scrollspy({
    //  target: '#sideNav'
    // });

})(jQuery); // End of use strict



function reviewPost(index) {
    var reviewUrl = $("#news-url").val();
    var url = "/review/" +reviewUrl;
    var request_method = "POST";
    var reviewTitle = $("#review-title").val();
    var reviewContent = $("#review-content").val();



    $.ajax({
        url: url,
        type: request_method,
        data:JSON.stringify({"reviewTitle" : reviewTitle, "reviewContent" : reviewContent, "newsUrl" : reviewUrl}),
    }).done(function (response) {
        window.location.href = response;
    });
}

function reviewDelete(index) {
    var url = $("#delete-form" + index).attr("action");


    var request_method = "DELETE";


    $.ajax({
        url: url,
        type: request_method,
    }).done(function (response) {
        window.location.href = response;
    });
}

function reviewModify() {
    var url = "/review/" +$("#news-url").val();
    var reviewTitle = $("#review-title").val();
    var reviewContent = $("#review-content").val();
    var request_method = "PUT";


    $.ajax({
        url: url,
        type: request_method,
        data:JSON.stringify({"reviewTitle" : reviewTitle, "reviewContent" : reviewContent}),
    }).done(function (response) {
        window.location.href = response;
    });
}