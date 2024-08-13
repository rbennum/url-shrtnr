document.addEventListener("DOMContentLoaded", function() {
    onFormSubmit()
});

function onFormSubmit() {
    const form = document.getElementById("url-form")
    form.addEventListener("submit", function(event) {
        event.preventDefault()
        const url = document.getElementById("url").value;
        const formData = {
            url: url
        }
        fetch("/url", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        })
        .then(response => response.json())
        .then(data => {
            const tag = data.short_tag;
            document.getElementById('content-view').style.display = 'block';
            document.getElementById('error-view').style.display = 'none';
            const contentText = document.getElementById('content-text')
            contentText.value = `localhost:8088/${tag}`;
            form.reset();
        })
        .catch(error => {
            document.getElementById('content-view').style.display = 'none';
            document.getElementById('error-view').style.display = 'block';
            const errorText = document.getElementById('error-text')
            errorText.value = `Error: ${error}`
            form.reset();
        })
    });
}