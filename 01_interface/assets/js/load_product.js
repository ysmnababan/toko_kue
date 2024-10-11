document.addEventListener('DOMContentLoaded', function() {

  const API_URL = 'http://localhost:3000'

    // Fetch and display user data
    fetch(`${API_URL}/products`) // Replace with your API URL
      .then(response => response.json()) // Parse JSON response
      .then(data => {
        console.log(data.data)
        const tbody = document.querySelector('#productTable tbody');
        tbody.innerHTML = ''; // Clear existing rows
        var count = 0;
        data.data.forEach(cat => {
            console.log(cat)
            count++
          const row = document.createElement('tr');
          row.setAttribute('data-prod-id', cat.id);

          row.innerHTML = `
            <td>${count}</td>
            <td class="id">${cat.id}</td>
            <td class="category_id">${cat.category_id}</td>
            <td class="name">${cat.name}</td>
            <td class="code">${cat.code}</td>
            <td class="price">${cat.price}</td>
            <td class="stock">${cat.stock}</td>
            <td class="actions">
              <button class="btn edit-btn" onclick="editProduct(${cat.id})">Edit</button>
              <button class="btn delete-btn" onclick="deleteProduct('${cat.id}')">Delete</button>
            </td>
          `;
  
          tbody.appendChild(row);
        });
      })
      .catch(error => {
        console.error('Error fetching user data:', error);
      });
  });
  
  // Function to handle editing a user
  function editProduct(id) {
    const row = document.querySelector(`tr[data-prod-id='${id}']`);
    const cidCell = row.querySelector('.category_id');
    const nameCell = row.querySelector('.name');
    const codeCell = row.querySelector('.code');
    const priceCell = row.querySelector('.price');
    const stockCell = row.querySelector('.stock');
    const actionsCell = row.querySelector('.actions');
  
    const currentcid = cidCell.textContent;
    const currentname = nameCell.textContent;
    const currentcode = codeCell.textContent;
    const currentprice = priceCell.textContent;
    const currentstock = stockCell.textContent;
  
    cidCell.innerHTML = `<input type="number" class="form-control" value="${currentcid}" required>`;
    nameCell.innerHTML = `<input type="text" class="form-control" value="${currentname}" required>`;
    codeCell.innerHTML = `<input type="text" class="form-control" value="${currentcode}" placeholder="" required>`;
    priceCell.innerHTML = `<input type="number" class="form-control" value="${currentprice}" placeholder="" required>`;
    stockCell.innerHTML = `<input type="number" class="form-control" value="${currentstock}" placeholder="" required>`;
    actionsCell.innerHTML = `
    <button class="btn save-btn" onclick="saveProduct(${id})">Save</button>
    <button class="btn cancel-btn" onclick="cancelEdit(${id})">Cancel</button>
  `;
  }
  
  function cancelEdit(id) {
    // Refresh the user data to cancel the edit
    location.reload();
  }
  
  function saveProduct(id) {
    const row = document.querySelector(`tr[data-prod-id='${id}']`);
  const cidInput = row.querySelector('.category_id input');
  const nameInput = row.querySelector('.name input');
  const codeInput = row.querySelector('.code input');
  const priceInput = row.querySelector('.price input');
  const stockInput = row.querySelector('.stock input');
  const errorMessage = document.querySelector('.error-message');

  const updatedcid = cidInput.value.trim();
  const updatedUser = nameInput.value.trim();
  const updatedCode = codeInput.value.trim();
  const updatedPrice = priceInput.value.trim();
  const updatedStock = stockInput.value.trim();

  // Validate that neither the username nor the password is empty
  if (!updatedUser || !updatedCode) {
    errorMessage.textContent = 'user or code cannot be empty.';
    errorMessage.style.display = 'block'; // Show the error message
    return; // Prevent saving if the username is empty
  } else {
    errorMessage.style.display = 'none'; // Hide the error message if both fields are valid
  }
  
    const formData = {
      category_id: parseInt(updatedcid),
      name: updatedUser,
      code: updatedCode,
      price:parseInt(updatedPrice),
      stock:parseInt(updatedStock)
    };
    console.log("HEEEREREE")
    console.log(formData)
  
    fetch(`${API_URL}/products/${id}`, {
      method: 'PUT', 
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    })
    .then(response => response.json())
    .then(result => {
      console.log('Product updated:', result);
      if (result.message === "Product updated successfully") {
        location.reload(); // Refresh the table after updating
      } else {
        console.error('Error updating Product:', result);
      }
    })
    .catch(error => {
      console.error('Error updating Product:', error);
    });
  }
  
  // Function to handle deleting a user
  function deleteProduct(id) {
    console.log(id)
    if (confirm('Are you sure you want to delete this ID?')) {

      fetch(`${API_URL}/products/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
          },
      })
      .then(response => response.json())
      .then(result => {
          console.log('User deleted:', result);
        if (result.message == "Product deleted successfully"){
            // Refresh the table after deletion
            location.reload();
        } else {
        console.error('Error deleting product:', result);
        }
      })
      .catch(error => {
        console.error('Error deleting product:', error);
      });
    }
  }
  