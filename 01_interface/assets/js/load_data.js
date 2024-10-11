document.addEventListener('DOMContentLoaded', function() {

  const API_URL = 'http://localhost:3000'

    // Fetch and display user data
    fetch(`${API_URL}/categories`) // Replace with your API URL
      .then(response => response.json()) // Parse JSON response
      .then(data => {
        console.log(data.data)
        const tbody = document.querySelector('#categoryTable tbody');
        tbody.innerHTML = ''; // Clear existing rows
        var count = 0;
        data.data.forEach(cat => {
            console.log(cat)
            count++
          const row = document.createElement('tr');
          row.setAttribute('data-cat-id', cat.id);

          row.innerHTML = `
            <td>${count}</td>
            <td class="id">${cat.id}</td>
            <td class="name">${cat.name}</td>
            <td class="code">${cat.code}</td>
            <td class="description">${cat.description}</td>
            <td class="actions">
              <button class="btn edit-btn" onclick="editCategory(${cat.id})">Edit</button>
              <button class="btn delete-btn" onclick="deleteCategory('${cat.id}')">Delete</button>
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
  function editCategory(id) {
    const row = document.querySelector(`tr[data-cat-id='${id}']`);
    const nameCell = row.querySelector('.name');
    const codeCell = row.querySelector('.code');
    const descriptionCell = row.querySelector('.description');
    const actionsCell = row.querySelector('.actions');
  
    const currentname = nameCell.textContent;
    const currentcode = codeCell.textContent;
    const currentdesc = descriptionCell.textContent;
  
    nameCell.innerHTML = `<input type="text" class="form-control" value="${currentname}" required>`;
    codeCell.innerHTML = `<input type="text" class="form-control" value="${currentcode}" placeholder="" required>`;
    descriptionCell.innerHTML = `<input type="text" class="form-control" value="${currentdesc}" placeholder="" required>`;
    actionsCell.innerHTML = `
    <button class="btn save-btn" onclick="saveCategory(${id})">Save</button>
    <button class="btn cancel-btn" onclick="cancelEdit(${id})">Cancel</button>
  `;
  }
  
  function cancelEdit(userId) {
    // Refresh the user data to cancel the edit
    location.reload();
  }
  
  function saveCategory(id) {
    const row = document.querySelector(`tr[data-cat-id='${id}']`);
  const nameInput = row.querySelector('.name input');
  const codeInput = row.querySelector('.code input');
  const descInput = row.querySelector('.description input');
  const errorMessage = document.querySelector('.error-message');

  const updatedUser = nameInput.value.trim();
  const updatedCode = codeInput.value.trim();
  const updatedDesc = descInput.value.trim();

  // Validate that neither the username nor the password is empty
  if (!updatedUser || !updatedCode) {
    errorMessage.textContent = 'user or code cannot be empty.';
    errorMessage.style.display = 'block'; // Show the error message
    return; // Prevent saving if the username is empty
  } else {
    errorMessage.style.display = 'none'; // Hide the error message if both fields are valid
  }
  
    const formData = {
      name: updatedUser,
      code: updatedCode,
      description:updatedDesc
    };
  
    fetch(`${API_URL}/categories/${id}`, {
      method: 'PUT', 
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    })
    .then(response => response.json())
    .then(result => {
      console.log('Category updated:', result);
      if (result.message === "Categories updated successfully") {
        location.reload(); // Refresh the table after updating
      } else {
        console.error('Error updating category:', result);
      }
    })
    .catch(error => {
      console.error('Error updating category:', error);
    });
  }
  
  // Function to handle deleting a user
  function deleteCategory(id) {
    console.log(id)
    if (confirm('Are you sure you want to delete this ID?')) {

      fetch(`${API_URL}/categories/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
          },
      })
      .then(response => response.json())
      .then(result => {
          console.log('User deleted:', result);
        if (result.message == "Categories deleted successfully"){
            // Refresh the table after deletion
            location.reload();
        } else {
        console.error('Error deleting category:', result);
        }
      })
      .catch(error => {
        console.error('Error deleting category:', error);
      });
    }
  }
  