document.getElementById('test').addEventListener('submit', function(e) {
    // e.preventDefault(); // Prevent the form from refreshing the page

    // Create a FormData object
    const formData = new FormData(this);
  
    // Log form data to see the input values
    console.log('Form data submitted:');
    for (let [key, value] of formData.entries()) {
      console.log(`${key}: ${value}`);
    }
  
    // Convert FormData to JSON for sending as the request body
    const formDataJson = Object.fromEntries(formData.entries());

    // Send the form data using fetch()
    fetch('http://localhost:3000/categories', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formDataJson), // Send JSON data
    })
    .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json(); // Parse the response as JSON
      })
      .then(data => {
        console.log('Success:', data);
        // Show the response to the user (e.g., in an alert or on the page)
        alert('Form submitted successfully: ' + JSON.stringify(data));
      })
      .catch((error) => {
        console.error('Error:', error);
        alert('There was a problem with the submission: ' + error.message);
      });
});

  