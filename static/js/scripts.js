    // Handle form submission to support multiline input
    document.getElementById('asciiForm').onsubmit = function(e) {
        e.preventDefault();
        const text = document.getElementById('text').value;
        const banner = document.getElementById('banner').value;
        fetch('/ascii-art', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams({
                text: text,
                banner: banner,
            }),
        })
        .then(response => response.text())
        .then(data => {
            document.open();
            document.write(data);
            document.close();
        });
    };