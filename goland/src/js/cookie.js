let url = new URL(window.location.href);
let params = url.searchParams;
let click_id = params.get('click_id');
let advertise_id = params.get('advertise_id');
document.cookie = `click_id=${click_id}`;
document.cookie = `advertise_id=${advertise_id}`;
