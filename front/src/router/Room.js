import React, { Component } from "react";
import styles from "./css/redactor.module.css"
//import { elements, elementGroups } from "../util/elements"
import Draggable, {DraggableCore} from 'react-draggable';
import { uuid } from "uuidv4"
import {Button} from "@design-system-rt/rtk-ui-kit"
import { ThemeProvider } from '@design-system-rt/rtk-ui-kit';
import { elements, elementGroups } from "../util/elements" 

class Room extends Component {
    constructor (props) {
      super(props)
      this.state = {
        ws : null,
        roomId : null,
        elements : [
            {
                id:uuid(),
                type : Button,
            }
        ],
        addItem : (el) => {
            console.log(this.state)
            this.state.elements.push({
                id : uuid(),
                type : el,
            })
            this.forceUpdate()
        },
        deleteItem : function(itemId) {
            this.elements = this.elements.filter(function(item) {
                return item !== itemId
            })
        }
    }
    }
    componentDidMount() {
      const roomId = this.props.match.params.roomId
      var ws = new WebSocket(`ws://la4z.xyz:8080/api/rooms/${roomId}?token=${localStorage.getItem("token")}`);
      ws.onopen = () => {
        console.log('connected')
      }
      ws.onmessage = evt => {
        const message = JSON.parse(evt.data)
        console.log(message)
      }
      ws.onclose = () => {
        console.log('disconnected')
      }
      this.setState({
        ws:ws,
        roomId : roomId
      })
      console.log(this.state)
    }
    render() {
      return (
        <ThemeProvider themeName="light">
        <main>
            <div className={styles.graph}>
                <div className={styles.viewport}>
                    <div className={styles.workplace}>
                    {
                        this.state.elements.map((el) => {
                            return (
                                <div>
                                    <Draggable key={el.id}>
                                        <div>
                                            {React.createElement(el.type)}
                                            <div onClick={() => this.setState({elements : this.state.elements.filter((item) => {
                                                return item.id !== el.id 
                                            })})}>
                                                Удалить
                                            </div>
                                        </div>
                                    </Draggable>
                                </div>
                            )
                        })
                    }
                    </div>
                </div>
                <aside>
                    <h3>Элементы</h3>
                    <div className={styles.acorcontainer}>
                        <input type="radio" name="acor" id="acor1" />
                        <label for="acor1">Холсты</label>
                        <div className={styles.acorbody}>

                        </div>
                        <input type="radio" name="acor" id="acor2" />
                        <label for="acor2">Шрифт</label>
                        <div className={styles.acorbody}>
                            <p>Описание вкладки</p>
                        </div>
                        <input type="radio" name="acor" id="acor3" />
                        <label for="acor3">Кнопки</label>
                        <div className={styles.acorbody}>
                        {
                            elementGroups.buttons.map((el) => {
                                return (
                                    <div onClick={() => {
                                        this.state.addItem(el)
                                        console.log(this.state.elements)
                                    }}>
                                        {React.createElement(el, {
                                            value : "Кнопка"
                                        })}
                                    </div>
                                )
                            })
                        }
                        </div>
                        <input type="radio" name="acor" id="acor4" />
                        <label for="acor4">Множественный выбор</label>
                        <div className={styles.acorbody}>
                            <p>Описание вкладки</p>
                        </div>
                        <input type="radio" name="acor" id="acor5" />
                        <label for="acor5">Иконки</label>
                        <div className={styles.acorbody}>
                        {
                            elementGroups.icon.map((el) => {
                                return (
                                    <div onClick={() => {
                                        this.state.addItem(el)
                                        console.log(this.state.elements)
                                    }}>
                                        {React.createElement(el, {
                                            
                                        })}
                                    </div>
                                )
                            })
                        }
                        </div>
                        <input type="radio" name="acor" id="acor6" />
                        <label for="acor6">Чекбоксы</label>
                        <div className={styles.acorbody}>
                        {
                            elementGroups.checkBoxes.map((el) => {
                                return (
                                    <div onClick={() => {
                                        this.state.addItem(el)
                                        console.log(this.state.elements)
                                    }}>
                                        {React.createElement(el, {
                                            
                                        })}
                                    </div>
                                )
                            })
                        }
                        </div>
                        <input type="radio" name="acor" id="acor7" />
                        <label for="acor7">Радио-кнопки</label>
                        <div className={styles.acorbody}>
                        {
                            elementGroups.radio.map((el) => {
                                return (
                                    <div onClick={() => {
                                        this.state.addItem(el)
                                        console.log(this.state.elements)
                                    }}>
                                        {React.createElement(el, {
                                            
                                        })}
                                    </div>
                                )
                            })
                        }
                        </div>
                        <input type="radio" name="acor" id="acor8" />
                        <label for="acor8">Переключатель</label>
                        <div className={styles.acorbody}>
                        {
                            elementGroups.switches.map((el) => {
                                return (
                                    <div onClick={() => {
                                        this.state.addItem(el)
                                        console.log(this.state.elements)
                                    }}>
                                        {React.createElement(el, {
                                            
                                        })}
                                    </div>
                                )
                            })
                        }
                        </div>
                        <input type="radio" name="acor" id="acor9" />
                        <label for="acor9">Табуляция</label>
                        <div className={styles.acorbody}>
                            <p>Описание вкладки</p>
                        </div>
                    </div>
                </aside>
            </div>
        </main>
        </ThemeProvider>
      )
    }
  }
  
  export default Room;