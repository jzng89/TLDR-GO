/* Set the width of the side navigation to 250px and the left margin of the page content to 250px */
var SideNavTog = true;

function toggleSideNav(){
  if (SideNavTog) {
    document.getElementById("leftsidenav").style.width = "0";
    document.getElementById("mainContent").style.marginLeft = "0";
    SideNavTog = false;
  }else{
    document.getElementById("leftsidenav").style.width = "250px";
    document.getElementById("mainContent").style.marginLeft = "280px";
    SideNavTog = true;
  }
}

/*when user scrolls down, shrink banner, sticky navigation bar*/
window.onscroll = function() {scrollFunction()};

// Get the navbar
var topnav = document.getElementById("topnav");
var navdrop = document.getElementById("dropdown-content");
var sidenav = document.getElementById("leftsidenav");
var header = document.getElementById("mainheader");
// Get the offset position of the navbar
var sticky = topnav.offsetTop;

function scrollFunction() {
  // Add the sticky class to the navbar when you reach its scroll position. Remove "sticky" when you leave the scroll position
  if (window.pageYOffset >= sticky) {
    topnav.classList.add("stickyTopNav");
    topnav.classList.add("stickyTopNav");
    navdrop.classList.add("stickydropdown");
    navdrop.style.position = "fixed";
    sidenav.classList.add("notransition");
    sidenav.style.position = "fixed";
    sidenav.style.top = "-230px";
    sidenav.offsetHeight; // Trigger a reflow, flushing the CSS changes
    sidenav.classList.remove("notransition");
  } else {
    topnav.classList.remove("stickyTopNav");
    navdrop.classList.remove("stickydropdown");
    navdrop.style.position = "absolute";
    sidenav.classList.add("notransition");
    sidenav.style.transition = "none !important";
    sidenav.style.position = "absolute";
    sidenav.style.top = "0";
    sidenav.offsetHeight; // Trigger a reflow, flushing the CSS changes
    sidenav.classList.remove("notransition");
  }

  /* this code is under consideration for minimized header
  if (document.body.scrollTop > 50 || document.documentElement.scrollTop > 50) {
      header.classList.add("stickyHeader");
    } else {
      header.classList.remove("stickyHeader");
    }
  */
}
