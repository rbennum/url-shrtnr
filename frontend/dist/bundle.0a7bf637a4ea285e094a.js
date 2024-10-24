/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./src/css/styles.css":
/*!****************************!*\
  !*** ./src/css/styles.css ***!
  \****************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId](module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/************************************************************************/
var __webpack_exports__ = {};
/*!*************************!*\
  !*** ./src/js/index.js ***!
  \*************************/
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _css_styles_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../css/styles.css */ "./src/css/styles.css");


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

    // try {
    //     const apiUrl = process.env.API_URL;
    //     console.log(apiUrl)
    //     const response = await fetch(apiUrl + '/url', {
    //         method: 'POST',
    //         headers: {
    //             'Content-Type': 'application/json'
    //         },
    //         body: JSON.stringify(formData)
    //     });

    //     if (!response.ok) {
    //         throw new Error('Unable to create new link. Please try again later.');
    //     }

    //     const data = await response.json();
    //     console.log(`Data: ${JSON.stringify(data)}`);
    //     const tag = data.tag;

    //     successMessage.classList.remove('hidden');
    //     errorView.classList.add('hidden');

    //     successMessage.textContent = 'URL: ';
    //     successMessage.innerHTML += `<a href="http://${tag}" target="_blank">${tag}</a>`;
    // } catch (error) {
    //     successMessage.classList.add('hidden');
    //     errorView.classList.remove('hidden');

    //     errorView.textContent = `${error}`;
    // }

    successMessage.classList.remove('hidden');
    errorView.classList.add('hidden');
    successMessage.textContent = 'URL: ';
    successMessage.innerHTML += `<a href="http://hello.com" target="_blank">hello.com</a>`;

    form.reset();
};

/******/ })()
;
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYnVuZGxlLjBhN2JmNjM3YTRlYTI4NWUwOTRhLmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7O0FBQUE7Ozs7Ozs7VUNBQTtVQUNBOztVQUVBO1VBQ0E7VUFDQTtVQUNBO1VBQ0E7VUFDQTtVQUNBO1VBQ0E7VUFDQTtVQUNBO1VBQ0E7VUFDQTtVQUNBOztVQUVBO1VBQ0E7O1VBRUE7VUFDQTtVQUNBOzs7OztXQ3RCQTtXQUNBO1dBQ0E7V0FDQSx1REFBdUQsaUJBQWlCO1dBQ3hFO1dBQ0EsZ0RBQWdELGFBQWE7V0FDN0Q7Ozs7Ozs7Ozs7QUNOMkI7O0FBRTNCO0FBQ0E7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSxnQkFBZ0I7QUFDaEI7QUFDQSxZQUFZOztBQUVaO0FBQ0E7QUFDQTs7QUFFQTtBQUNBLGdDQUFnQyxxQkFBcUI7QUFDckQ7O0FBRUE7QUFDQTs7QUFFQTtBQUNBLDBEQUEwRCxJQUFJLG9CQUFvQixJQUFJO0FBQ3RGLFNBQVM7QUFDVDtBQUNBOztBQUVBLHNDQUFzQyxNQUFNO0FBQzVDOztBQUVBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0EiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly91cmwtc2hydG5yX2ZlLy4vc3JjL2Nzcy9zdHlsZXMuY3NzP2U0YjAiLCJ3ZWJwYWNrOi8vdXJsLXNocnRucl9mZS93ZWJwYWNrL2Jvb3RzdHJhcCIsIndlYnBhY2s6Ly91cmwtc2hydG5yX2ZlL3dlYnBhY2svcnVudGltZS9tYWtlIG5hbWVzcGFjZSBvYmplY3QiLCJ3ZWJwYWNrOi8vdXJsLXNocnRucl9mZS8uL3NyYy9qcy9pbmRleC5qcyJdLCJzb3VyY2VzQ29udGVudCI6WyIvLyBleHRyYWN0ZWQgYnkgbWluaS1jc3MtZXh0cmFjdC1wbHVnaW5cbmV4cG9ydCB7fTsiLCIvLyBUaGUgbW9kdWxlIGNhY2hlXG52YXIgX193ZWJwYWNrX21vZHVsZV9jYWNoZV9fID0ge307XG5cbi8vIFRoZSByZXF1aXJlIGZ1bmN0aW9uXG5mdW5jdGlvbiBfX3dlYnBhY2tfcmVxdWlyZV9fKG1vZHVsZUlkKSB7XG5cdC8vIENoZWNrIGlmIG1vZHVsZSBpcyBpbiBjYWNoZVxuXHR2YXIgY2FjaGVkTW9kdWxlID0gX193ZWJwYWNrX21vZHVsZV9jYWNoZV9fW21vZHVsZUlkXTtcblx0aWYgKGNhY2hlZE1vZHVsZSAhPT0gdW5kZWZpbmVkKSB7XG5cdFx0cmV0dXJuIGNhY2hlZE1vZHVsZS5leHBvcnRzO1xuXHR9XG5cdC8vIENyZWF0ZSBhIG5ldyBtb2R1bGUgKGFuZCBwdXQgaXQgaW50byB0aGUgY2FjaGUpXG5cdHZhciBtb2R1bGUgPSBfX3dlYnBhY2tfbW9kdWxlX2NhY2hlX19bbW9kdWxlSWRdID0ge1xuXHRcdC8vIG5vIG1vZHVsZS5pZCBuZWVkZWRcblx0XHQvLyBubyBtb2R1bGUubG9hZGVkIG5lZWRlZFxuXHRcdGV4cG9ydHM6IHt9XG5cdH07XG5cblx0Ly8gRXhlY3V0ZSB0aGUgbW9kdWxlIGZ1bmN0aW9uXG5cdF9fd2VicGFja19tb2R1bGVzX19bbW9kdWxlSWRdKG1vZHVsZSwgbW9kdWxlLmV4cG9ydHMsIF9fd2VicGFja19yZXF1aXJlX18pO1xuXG5cdC8vIFJldHVybiB0aGUgZXhwb3J0cyBvZiB0aGUgbW9kdWxlXG5cdHJldHVybiBtb2R1bGUuZXhwb3J0cztcbn1cblxuIiwiLy8gZGVmaW5lIF9fZXNNb2R1bGUgb24gZXhwb3J0c1xuX193ZWJwYWNrX3JlcXVpcmVfXy5yID0gKGV4cG9ydHMpID0+IHtcblx0aWYodHlwZW9mIFN5bWJvbCAhPT0gJ3VuZGVmaW5lZCcgJiYgU3ltYm9sLnRvU3RyaW5nVGFnKSB7XG5cdFx0T2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIFN5bWJvbC50b1N0cmluZ1RhZywgeyB2YWx1ZTogJ01vZHVsZScgfSk7XG5cdH1cblx0T2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsICdfX2VzTW9kdWxlJywgeyB2YWx1ZTogdHJ1ZSB9KTtcbn07IiwiaW1wb3J0ICcuLi9jc3Mvc3R5bGVzLmNzcyc7XG5cbmNvbnN0IG9uU3VibWl0ID0gYXN5bmMgKGUpID0+IHtcbiAgICBjb25zb2xlLmxvZyhcIm9uU3VibWl0IGNhbGxlZFwiKTtcbiAgICBlLnByZXZlbnREZWZhdWx0KCk7XG5cbiAgICBjb25zdCBmb3JtID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ3VybEZvcm0nKTtcbiAgICBjb25zdCB1cmxJbnB1dCA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCd1cmxJbnB1dCcpO1xuICAgIGNvbnN0IGVycm9yTWVzc2FnZSA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdlcnJvck1lc3NhZ2UnKTtcbiAgICBjb25zdCByZXN1bHREaXYgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgncmVzdWx0Jyk7XG4gICAgY29uc3Qgc3VjY2Vzc01lc3NhZ2UgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnc3VjY2Vzc01lc3NhZ2UnKTtcbiAgICBjb25zdCBlcnJvclZpZXcgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnZXJyb3JWaWV3Jyk7XG5cbiAgICBpZiAodXJsSW5wdXQudmFsdWUubGVuZ3RoID4gNDAwKSB7XG4gICAgICAgIGVycm9yTWVzc2FnZS5jbGFzc0xpc3QucmVtb3ZlKCdoaWRkZW4nKTtcbiAgICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIHJlc3VsdERpdi5jbGFzc0xpc3QucmVtb3ZlKCdoaWRkZW4nKTtcbiAgICBlcnJvck1lc3NhZ2UuY2xhc3NMaXN0LmFkZCgnaGlkZGVuJyk7XG5cbiAgICBjb25zdCBmb3JtRGF0YSA9IHtcbiAgICAgICAgb3JpZ2luYWxfdXJsOiB1cmxJbnB1dC52YWx1ZVxuICAgIH1cblxuICAgIC8vIHRyeSB7XG4gICAgLy8gICAgIGNvbnN0IGFwaVVybCA9IHByb2Nlc3MuZW52LkFQSV9VUkw7XG4gICAgLy8gICAgIGNvbnNvbGUubG9nKGFwaVVybClcbiAgICAvLyAgICAgY29uc3QgcmVzcG9uc2UgPSBhd2FpdCBmZXRjaChhcGlVcmwgKyAnL3VybCcsIHtcbiAgICAvLyAgICAgICAgIG1ldGhvZDogJ1BPU1QnLFxuICAgIC8vICAgICAgICAgaGVhZGVyczoge1xuICAgIC8vICAgICAgICAgICAgICdDb250ZW50LVR5cGUnOiAnYXBwbGljYXRpb24vanNvbidcbiAgICAvLyAgICAgICAgIH0sXG4gICAgLy8gICAgICAgICBib2R5OiBKU09OLnN0cmluZ2lmeShmb3JtRGF0YSlcbiAgICAvLyAgICAgfSk7XG5cbiAgICAvLyAgICAgaWYgKCFyZXNwb25zZS5vaykge1xuICAgIC8vICAgICAgICAgdGhyb3cgbmV3IEVycm9yKCdVbmFibGUgdG8gY3JlYXRlIG5ldyBsaW5rLiBQbGVhc2UgdHJ5IGFnYWluIGxhdGVyLicpO1xuICAgIC8vICAgICB9XG5cbiAgICAvLyAgICAgY29uc3QgZGF0YSA9IGF3YWl0IHJlc3BvbnNlLmpzb24oKTtcbiAgICAvLyAgICAgY29uc29sZS5sb2coYERhdGE6ICR7SlNPTi5zdHJpbmdpZnkoZGF0YSl9YCk7XG4gICAgLy8gICAgIGNvbnN0IHRhZyA9IGRhdGEudGFnO1xuXG4gICAgLy8gICAgIHN1Y2Nlc3NNZXNzYWdlLmNsYXNzTGlzdC5yZW1vdmUoJ2hpZGRlbicpO1xuICAgIC8vICAgICBlcnJvclZpZXcuY2xhc3NMaXN0LmFkZCgnaGlkZGVuJyk7XG5cbiAgICAvLyAgICAgc3VjY2Vzc01lc3NhZ2UudGV4dENvbnRlbnQgPSAnVVJMOiAnO1xuICAgIC8vICAgICBzdWNjZXNzTWVzc2FnZS5pbm5lckhUTUwgKz0gYDxhIGhyZWY9XCJodHRwOi8vJHt0YWd9XCIgdGFyZ2V0PVwiX2JsYW5rXCI+JHt0YWd9PC9hPmA7XG4gICAgLy8gfSBjYXRjaCAoZXJyb3IpIHtcbiAgICAvLyAgICAgc3VjY2Vzc01lc3NhZ2UuY2xhc3NMaXN0LmFkZCgnaGlkZGVuJyk7XG4gICAgLy8gICAgIGVycm9yVmlldy5jbGFzc0xpc3QucmVtb3ZlKCdoaWRkZW4nKTtcblxuICAgIC8vICAgICBlcnJvclZpZXcudGV4dENvbnRlbnQgPSBgJHtlcnJvcn1gO1xuICAgIC8vIH1cblxuICAgIHN1Y2Nlc3NNZXNzYWdlLmNsYXNzTGlzdC5yZW1vdmUoJ2hpZGRlbicpO1xuICAgIGVycm9yVmlldy5jbGFzc0xpc3QuYWRkKCdoaWRkZW4nKTtcbiAgICBzdWNjZXNzTWVzc2FnZS50ZXh0Q29udGVudCA9ICdVUkw6ICc7XG4gICAgc3VjY2Vzc01lc3NhZ2UuaW5uZXJIVE1MICs9IGA8YSBocmVmPVwiaHR0cDovL2hlbGxvLmNvbVwiIHRhcmdldD1cIl9ibGFua1wiPmhlbGxvLmNvbTwvYT5gO1xuXG4gICAgZm9ybS5yZXNldCgpO1xufTtcbiJdLCJuYW1lcyI6W10sInNvdXJjZVJvb3QiOiIifQ==