// Handle form submission to support multiline input

document.getElementById('asciiForm').onsubmit = function(e) {
    e.preventDefault(); // Prevent the default form submission behavior

    // Get text and banner values from form inputs
    const text = document.getElementById('text').value;
    const banner = document.getElementById('banner').value;

    // Send a POST request to server endpoint '/ascii-art'
    fetch('/ascii-art', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        // Send text and banner as URL-encoded parameters in the request body
        body: new URLSearchParams({
            text: text,
            banner: banner,
        }),
    })
    .then(response => response.text()) // Parse the response as text
    .then(data => {
        // Update the browser URL without reloading the page
        window.history.pushState({}, '', '/ascii-art');
        // Replace the current document content with the received ASCII art
        document.open();
        document.write(data);
        document.close();
    })
    .catch(error => {
        console.error('Error:', error); // Log any errors to the console
    });
};

// hamburger menu
document.addEventListener('DOMContentLoaded', () => {
    const hamburger = document.querySelector('.fa-bars');
    const navLinks = document.querySelector('.nav-links');

    hamburger.addEventListener('click', () => {
        navLinks.classList.toggle('active');
    });
});
