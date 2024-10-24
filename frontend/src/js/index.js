import '../css/styles.css';

// Declare form handler function first
const onSubmit = async (e) => {
    console.log("onSubmit called");
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
        original_url: urlInput.value
    }

    try {
        const apiUrl = process.env.API_URL;
        console.log(apiUrl)
        const response = await fetch(apiUrl + '/url', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });

        if (!response.ok) {
            throw new Error('Unable to create new link. Please try again later.');
        }

        const response_data = await response.json();
        console.log(`Data: ${JSON.stringify(response_data)}`);
        let data = response_data.data
        let short_url = data.short_url

        successMessage.classList.remove('hidden');
        errorView.classList.add('hidden');

        successMessage.textContent = 'URL: ';
        successMessage.innerHTML += `<a href="${short_url}" target="_blank">${short_url}</a>`;
    } catch (error) {
        successMessage.classList.add('hidden');
        errorView.classList.remove('hidden');
        errorView.textContent = `${error}`;
    }

    form.reset();
};

// Add the event listener when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('urlForm');
    form.addEventListener('submit', onSubmit);
});

// Make onSubmit available globally
window.onSubmit = onSubmit;

// Export for module usage
export { onSubmit };