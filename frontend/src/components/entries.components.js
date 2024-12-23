import React, { useState, useEffect } from "react";
import axios from "axios";
import { Button, Form, Container, Modal } from "react-bootstrap";

import Entry from "./single-entry.components";

export default function Entries() {
  const [entries, setEntries] = useState([]);
  const [refreshData, setRefreshData] = useState(false);
  const [changeEntry, setChangeEntry] = useState({ change: false, id: 0 });
  const [changeIngredient, setChangeIngredient] = useState({
    change: false,
    id: 0,
  });
  const [newIngredientName, setNewIngredientName] = useState("");
  const [addNewEntry, setAddNewEntry] = useState(false);
  const [newEntry, setNewEntry] = useState({
    dish: "",
    ingredients: "",
    calories: 0,
    fats: 0,
  });

  const getAllEntries = () => {
    var url = "http://localhost:8000/entries";
    axios.get(url, { responseType: "json" }).then((res) => {
      if (res.status === 200) {
        setEntries(res.data);
      }
    });
  };

  const addSingleEntry = () => {
    setAddNewEntry(false);
    var url = "http://localhost:8000/entry/create";
    axios
      .post(url, {
        ingredients: newEntry.ingredients,
        fats: parseFloat(newEntry.fats),
        dish: newEntry.dish,
        calories: newEntry.calories,
      })
      .then((response) => {
        if (response.status === 200) {
          setRefreshData(true);
        }
      });
  };
  const changeSingleEntry = () => {
    changeEntry.change = false;
    var url = "http://localhost:8000/entry/update/" + changeEntry.id;
    axios.put(url, newEntry).then((res) => {
      if (res.status === 200) {
        setRefreshData(true);
      }
    });
  };
  const changeIngredientForEntry = () => {
    changeIngredient.change = false;
    var url = "http://localhost:8000/ingredient/update/" + changeIngredient.id;
    axios
      .put(url, {
        ingredients: newIngredientName,
      })
      .then((res) => {
        console.log(res.status);
        if (res.status === 200) {
          setRefreshData(true);
        }
      });
  };

  const deleteSingleEntry = (id) => {
    var url = "http://localhost:8000/entry/delete/" + id;
    axios.delete(url, {}).then((response) => {
      if (response.status === 200) {
        setRefreshData(true);
      }
    });
  };

  useEffect(() => {
    getAllEntries();
  }, []);

  if (refreshData) {
    setRefreshData(false);
    getAllEntries();
  }
  return (
    <div>
      <Container>
        <Button onClick={() => setAddNewEntry(true)}>
          Track today's calories
        </Button>
      </Container>
      <Container>
        {entries !== null &&
          entries.map((entry) => (
            <Entry
              entryData={entry}
              deleteSingleEntry={deleteSingleEntry}
              setChangeIngredient={setChangeIngredient}
              setChangeEntry={setChangeEntry}
            />
          ))}
      </Container>
      {/* modal for adding single entry */}
      <Modal show={addNewEntry} onHide={() => setAddNewEntry(false)} centered>
        <Modal.Header closeButton>
          <Modal.Title>Add Calorie Entry</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Dish</Form.Label>
            <Form.Control onChange={(e) => (newEntry.dish = e.target.value)} />
            <Form.Label>Ingredients</Form.Label>
            <Form.Control
              onChange={(e) => (newEntry.ingredients = e.target.value)}
            />
            <Form.Label>Calories</Form.Label>
            <Form.Control
              onChange={(e) => (newEntry.calories = e.target.value)}
            />
            <Form.Label>Fats</Form.Label>
            <Form.Control
              type="number"
              onChange={(e) => (newEntry.fats = e.target.value)}
            />
          </Form.Group>
          <Button onClick={addSingleEntry}>Add</Button>
          <Button onClick={() => setAddNewEntry(false)}>Cancel</Button>
        </Modal.Body>
      </Modal>
      {/* modal for changing the ingredient only */}
      <Modal
        show={changeIngredient.change}
        onHide={() => setChangeIngredient({ change: false, id: 0 })}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Change Ingredient</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>New Ingredients</Form.Label>
            <Form.Control
              onChange={(e) => setNewIngredientName(e.target.value)}
            />
          </Form.Group>
          <Button onClick={changeIngredientForEntry}>Change</Button>
          <Button onClick={() => setChangeIngredient({ change: false, id: 0 })}>
            Cancel
          </Button>
        </Modal.Body>
      </Modal>
      {/* modal for changing entire entry */}
      <Modal
        show={changeEntry.change}
        onHide={() => setChangeEntry({ change: false, id: 0 })}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Change Entry</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Dish</Form.Label>
            <Form.Control
              onChange={(e) => (newEntry.dish = e.target.value)}
              value={!newEntry.dish && newEntry.dish}
            />
            <Form.Label>Ingredients</Form.Label>
            <Form.Control
              onChange={(e) => (newEntry.ingredients = e.target.value)}
              value={!newEntry.ingredients && newEntry.ingredients}
            />
            <Form.Label>Calories</Form.Label>
            <Form.Control
              onChange={(e) => (newEntry.calories = e.target.value)}
              value={!newEntry.calories && newEntry.calories}
            />
            <Form.Label>Fats</Form.Label>
            <Form.Control
              type="number"
              onChange={(e) => (newEntry.fats = e.target.value)}
              value={!newEntry.fats && newEntry.fats}
            />
          </Form.Group>
          <Button onClick={changeSingleEntry}>Update</Button>
          <Button onClick={() => setChangeEntry({ change: false, id: 0 })}>
            Cancel
          </Button>
        </Modal.Body>
      </Modal>
    </div>
  );
}
