const onSubmit = async (e) => {
    e.preventDefault();

    const form = document.getElementById('urlForm');
    const urlInput = document.getElementById('urlInput');
    const errorMessage = document.getElementById('errorMessage');
    const resultDiv = document.getElementById('result');
    const successMessage = document.getElementById('successMessage');
    const errorView = document.getElementById('errorView');

    if (urlInput.value.length > 400) {
        errorMessage.classList.remove('hidden');
        return;
    }

    resultDiv.classList.remove('hidden');
    errorMessage.classList.add('hidden');

    const formData = {
        url: urlInput.value
    }

    try {
        const response = await fetch('/url', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });

        if (!response.ok) {
            throw new Error('Unable to create new link. Please try again later.');
        }

        const data = await response.json();
        console.log(`Data: ${JSON.stringify(data)}`);
        const tag = data.tag;

        successMessage.classList.remove('hidden');
        errorView.classList.add('hidden');

        successMessage.textContent = 'URL: ';
        successMessage.innerHTML += `<a href="http://${tag}" target="_blank">${tag}</a>`;
    } catch (error) {
        successMessage.classList.add('hidden');
        errorView.classList.remove('hidden');

        errorView.textContent = `${error}`;
    }

    form.reset();
};
