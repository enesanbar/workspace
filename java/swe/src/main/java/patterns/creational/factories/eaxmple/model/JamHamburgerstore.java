package patterns.creational.factories.eaxmple.model;

import patterns.creational.factories.eaxmple.interfaces.HamburgerStore;

public class JamHamburgerstore extends HamburgerStore {
    @Override
    public Hamburger createHamburger(String type) {

        if (type.equals("cheese")) {
            return new JamaicanCheeseBurger();
        }else if (type.equals("veggie")) {
            return new JamaicanVeggieBurger();

        }else  return null;
    }
}
