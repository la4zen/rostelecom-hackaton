import * as elms from "@design-system-rt/rtk-ui-kit"
export const elements = {
    button : elms.Button,
    iconButton : elms.IconButton,
    radioButton : elms.RadioButton,
    functionButton : elms.FunctionButton,
    inputText : elms.InputText,
    inputAmount : elms.InputAmount,
    inputEmail : elms.InputEmail,
    inputNumberStepper : elms.InputNumberStepper,
    inputPhone : elms.InputPhone,
    multiSellect : elms.Multiselect,
    radioGroup : elms.RadioGroup,
    radioButton : elms.RadioButton,
    iconButton : elms.IconButton,
    checkBox : elms.Checkbox,
    switch : elms.Switch,
}

export const elementGroups = {
    buttons : [
        elements.button, elements.iconButton, elements.radioButton, elements.functionButton
    ],
    inputs : [
        elements.inputText, elements.inputAmount, elements.inputEmail, elements.inputNumberStepper, elements.inputPhone
    ],
    multiSellects : [
        elements.multiSellect
    ],
    radio : [
        elements.radioButton, elements.radioGroup
    ],
    icon : [
        elements.iconButton
    ],
    checkBoxes : [
        elements.checkBox
    ],
    switches : [
        elements.switch
    ]
}
