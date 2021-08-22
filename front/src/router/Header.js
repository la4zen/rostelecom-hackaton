import { Component } from "react";
import styles from "./css/redactor.module.css"
import logo from "./img/logo.png"

export default class Header extends Component {
    render() {
        return (
            <div className={styles.graphheader}>
                <div className={styles.headeritems}>
                    <a class={styles.headeritem} href="/">
                        <img src={logo} alt=""/>
                        <p>RT Graph</p>
                    </a>
                    <div className={styles.headeritem}>
                        <button type="button" className={styles.share}>Поделиться</button>
                        <button type="button" className={styles.round}></button>
                    </div>
                </div>
            </div>
        )
    }
}